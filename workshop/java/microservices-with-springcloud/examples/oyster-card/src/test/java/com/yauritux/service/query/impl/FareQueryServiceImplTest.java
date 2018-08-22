package com.yauritux.service.query.impl;

import static org.junit.Assert.assertTrue;
import static org.mockito.Mockito.when;

import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;

import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.context.junit4.SpringRunner;

import com.yauritux.model.constant.JourneyType;
import com.yauritux.model.entity.Fare;
import com.yauritux.repository.FareRepository;
import com.yauritux.service.query.FareQueryService;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
@RunWith(SpringRunner.class)
public class FareQueryServiceImplTest {

	private FareQueryService fareQueryService;
	
	@MockBean
	private FareRepository fareRepository;
	
	@Before
	public void setup() {
		this.fareQueryService = new FareQueryServiceImpl(fareRepository);
	}
	
	@Test
	public void findByJourneyType_existingFare_recordIsPresent() {
		Fare fare = new Fare();
		fare.setJourney("Bus Fare");
		fare.setJourneyType(JourneyType.BUS);
		fare.setFare(BigDecimal.valueOf(1.80));
		List<Fare> fares = new ArrayList<>();
		fares.add(fare);
		when(fareRepository.findByJourneyType(JourneyType.BUS)).thenReturn(fares);
		assertTrue(fareQueryService.findByJourneyType(JourneyType.BUS).isPresent());
	}	
}
