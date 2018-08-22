package com.yauritux.repository;

import static org.junit.Assert.assertNotNull;
import static org.junit.Assert.assertTrue;

import java.math.BigDecimal;
import java.util.List;
import java.util.Map;

import org.junit.After;
import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.orm.jpa.DataJpaTest;
import org.springframework.boot.test.autoconfigure.orm.jpa.TestEntityManager;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import com.yauritux.Application;
import com.yauritux.DataTestConfiguration;
import com.yauritux.model.constant.JourneyType;
import com.yauritux.model.entity.Fare;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
@RunWith(SpringRunner.class)
@SpringBootTest(classes = { Application.class, DataTestConfiguration.class })
@DataJpaTest
public class FareRepositoryTest {
	
	@Autowired
	private TestEntityManager em;
	
	@Autowired
	private FareRepository fareRepository;
	
	@Autowired
	private Map<String, Fare> fares;
		
	@Before
	public void setup() {
		Fare fare = fares.get("BusFare");
		em.merge(fare);
		em.flush();		
	}
	
	@After
	public void tearDown() {
		fareRepository.delete(fares.get("BusFare"));
	}
	
	@Test
	public void findByJourneyType_existingFare_fareRecordsNotNull() {		
		List<Fare> fareRecords = fareRepository.findByJourneyType(JourneyType.BUS);
		assertNotNull(fareRecords);
		assertTrue(fareRecords.size() > 0);
	}
	
	@Test
	public void findByJourneyType_existingFare_journeyTypeMatched() {
		List<Fare> fareRecords = fareRepository.findByJourneyType(JourneyType.BUS);
		assertTrue(fareRecords.get(0).getJourneyType() == JourneyType.BUS);
	}
	
	@Test
	public void findByJourneyType_existingFare_priceMoreThanZero() {
		List<Fare> fareRecords = fareRepository.findByJourneyType(JourneyType.BUS);
		assertTrue(fareRecords.get(0).getFare().compareTo(BigDecimal.ZERO) > 0);
	}
}
