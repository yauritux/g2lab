package com.yauritux.repository;

import static org.junit.Assert.assertTrue;

import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.List;
import java.util.stream.Collectors;

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
import com.yauritux.model.entity.Card;
import com.yauritux.model.entity.CardTransaction;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
@RunWith(SpringRunner.class)
@SpringBootTest(classes = { Application.class, DataTestConfiguration.class })
@DataJpaTest(showSql = true)
public class CardTransactionRepositoryTest {

	@Autowired
	private TestEntityManager em;
	
	@Autowired
	private CardTransactionRepository cardTransactionRepository;
	
	@Autowired
	private Card card;
	
	private CardTransaction cardTrx;
	
	@Test
	public void findByCard_hasSomeTransactions_someRecordsReturned() {
		em.persist(card);
		em.flush();
		
		cardTrx = new CardTransaction();
		cardTrx.setJourneyType(JourneyType.BUS);
		cardTrx.setCard(card);
		cardTrx.setDate(java.sql.Date.valueOf(LocalDateTime.now().toLocalDate()));
		em.persist(cardTrx);
		
		em.flush();				
		
		List<CardTransaction> cardTransactions = cardTransactionRepository.findByCard(card).stream().collect(Collectors.toList());
		assertTrue(cardTransactions.size() > 0);
	}
	
	@Test
	public void findByCard_noTransactions_noRecordsFound() {
		Card c = new Card();
		c.setSerialNo("555");
		c.setOwner("Khairunnisa");
		c.setBalance(BigDecimal.valueOf(500));
		em.persist(c);
		em.flush();
		List<CardTransaction> cardTransactions = cardTransactionRepository.findByCard(c).stream().collect(Collectors.toList());
		assertTrue(cardTransactions.size() == 0);
	}
	
	@Test
	public void findAll_noTransactions_noRecordsFound() {
		List<CardTransaction> cardTransactions = cardTransactionRepository.findAll();
		assertTrue(cardTransactions.size() == 0);
	}
}
